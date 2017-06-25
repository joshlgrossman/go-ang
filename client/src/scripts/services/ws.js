app.factory('ws', function(){

	function ws(path){
		return '/ws/' + path
	}

	ws.tasks = ws('tasks')

	return ws

})