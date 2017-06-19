app.component('tasks', {

	templateUrl: './tasks.view.html',

	controller: ['$http', 'ws', function($http, ws) {
		$http.get(ws('tasks')).then(response => {
			this.tasks = response.data
		})
	}]

})