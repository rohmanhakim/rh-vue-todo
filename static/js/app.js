Vue.component('confirmation-modal', {
    template: '#confirmation-modal-template',
    props: ['show', 'index','title', 'message','type','callback'],
    events: {
        onShowConfirmationModal: function(index,title,message,type,callback) {
            this.show = true;           // specify if this modal should be shown or not
            this.index = index;         // specify the related task's index
            this.title = title;         // specify this modal dialog's title
            this.message = message;     // specify the this modal dialog's message
            this.type = type;           // specify if this modal dialog type is positive or negative
            this.callback = callback;   // specify the callback for this modal dialog
        }
    },
    methods: {
        onCancel: function () {
            this.show = false;
        },
        onConfirm: function(index) {
            this.$dispatch(this.callback,index);
            console.log("Sending dispatch on confirm delete");
            this.show = false;
        }
    }
});

new Vue({
    // we want to target the div with an id of 'events'
    el: '#tasks',

    //here we can register any values or collections that hold data
    // for the application
    data: {
        task: { id: '', title: '', notes: '' },
        tasksContainer: { tasks: [] }
    },

    //anything within the ready function will run when the application loads
    ready: function() {
        // when the application loads, we want to call the method that initializes some data
        this.getAllTasks();
    },
    
    //anything within the events options will listen for the specified events that dispatched fro the child component/instance
    events: {
        //this will listen for 'onConfirmDeleteTask' from child who sends it
        'onConfirmDeleteTask' : function(index) {
            this.deleteTask(index);
        }
    },

    //methods we want to use in our application are registered here
    methods: {

        // we dedicate a method to retrieving and setting some data
        getAllTasks: function() {

            this.$http.get('http://localhost:8080/task/all').success(function(tasks) {

                // set the tasks array equals to the 'tasks' property of the response
                this.$set('tasksContainer', tasks);
            }).error(function(error) {
                console.log(error);
            });

        },

        // adds an event to the existing events array
        addTask: function(){
            if(this.task.title) {
                this.$http.post('http://localhost:8080/task', this.task).success(function(response) {

                    // set the task id equals to the 'id' property of the response
                    this.task.id = response.id

                    // push task to the array
                    this.tasksContainer.tasks.push(this.task);

                    // set the task to empty
                    this.task = { id: '', title: '', notes: '' };

                    console.log("Task added!");
                }).error(function(error) {
                    console.log(error);
                });
            }
        },

        //delete a task
        deleteTask: function (index) {
                console.log("receive dispatch event on confirm delete");
                this.$http.delete('http://localhost:8080/task/' + this.tasksContainer.tasks[index].id)
                .success( function(response) {
                // $remove is a Vue convenience method similar to splice
                    this.tasksContainer.tasks.$remove(this.tasksContainer.tasks[index]);
                })
                .error( function(error) {
                        console.log(error);
                });
        },

        showDeleteTaskModal: function (index) {
            //broadcast will send the specified event to all child of this vue's instance
            //here, we pass fouur argument to the receiving child: the index of the task, the modal dialog's message, the type of the modal dialog ("success" for positive, and "danger" for negative), and its callback wich will be listened by this instance's "event" options
            this.$broadcast('onShowConfirmationModal',index,"Confirm delete","Are you sure want to delete this task?","danger",'onConfirmDeleteTask');
        }
    }
});