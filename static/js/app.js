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

        deleteTask: function (index) {
            if(confirm("Are you sure you want to delete this task?")){

                this.$http.delete('http://localhost:8080/task/' + this.tasksContainer.tasks[index].id)
                    .success(function(response) {

                        // $remove is a Vue convenience method similar to splice
                        this.tasksContainer.tasks.$remove(this.tasksContainer.tasks[index]);
                    })
                    .error(function(error) {
                        console.log(error);
                    });
            }
        }
    }
});