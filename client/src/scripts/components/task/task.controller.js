app.component('task', {

	templateUrl: './task.view.html',

	bindings: {
		task: '<'
	},

	controller: ['$http', '$scope', 'ws', function($http, $scope, ws){

		this.editing = false;
		this.prev = angular.copy(this.task)

		this.edit = () => {
			this.prev = angular.copy(this.task)
			this.editing = true
		}

		this.cancel = () => {
			this.task = this.prev
			this.editing = false
		}

		this.update = () => {
			$http.put(ws.tasks, this.task).then(response => {

				if(!response.data.error)
					$scope.$emit('task:updated', this.task)

			}).catch(console.log)
			this.editing = false
		}

	}]

})