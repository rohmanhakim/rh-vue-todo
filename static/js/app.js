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
                this.$set('taskContainer', tasks);
            }).error(function(error) {
                console.log(error);
            });

            /*var tasks = [
                {
                    id: '1',
                    title: 'TIFF',
                    notes: 'Toronto International Film Festival'
                },
                {
                    id: '2',
                    title: 'The Martian Premiere',
                    notes: 'The Martian comes to theatres.'
                },
                {
                    id: '3',
                    title: 'SXSW',
                    notes: 'Music, film and interactive festival in Austin, TX.'
                }
            ];*/
            // $set is a convenience method provided by Vue that is similar to pushing
            // data onto an array
            //this.$set('tasks', tasks);

        },

        // adds an event to the existing events array
        addTask: function(){
            if(this.task.name) {
                this.taskContainer.tasks.push(this.task);
                this.task = { title: '', notes: ''};
            }
        },

        deleteTask: function (index) {
            if(confirm("Are you sure you want to delete this task?")){
                // $remove is a Vue convenience method similar to splice
                this.taskContainer.tasks.$remove(this.tasks[index]);
            }
        }
    }
});