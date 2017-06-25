app.component('taskForm', {

	templateUrl: './task-form.view.html',

	controller: ['$http', '$scope', 'ws', function($http, $scope, ws){
		this.task = {}

		this.reset = () => {
			this.task = {
				title: '',
				description: ''
			}
		}

		this.submit = () => {
			$http.post(ws.tasks, angular.copy(this.task)).then(response => {

				if(!response.data.error)
					$scope.$emit('task:added', response.data)

			}).catch(console.log)
			
			this.reset()
		}

		this.reset()

	}]

})