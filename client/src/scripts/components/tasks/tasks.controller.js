app.component('tasks', {

	templateUrl: './tasks.view.html',

	controller: ['$http', '$scope', 'ws', function($http, $scope, ws) {

		$http.get(ws('tasks')).then(response => {
			this.tasks = response.data
		})

		$scope.$on('task:added', (event, task) => {
			this.tasks.push(task)
		})

	}]

})