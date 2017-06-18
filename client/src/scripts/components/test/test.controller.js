app.component('test', {

  templateUrl: './test.view.html',

  controller: ['$http', 'ws', function($http, ws) {
    $http.get(ws('test')).then(r => {
      this.msg = r.data
    })
  }]

})