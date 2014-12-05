(function(){
	angular.module("godel").config(function($routeProvider) {
	  $routeProvider
	    .when('/', {
	      controller:'HomeController',
	      templateUrl:'tpl/home.html'
	    })
	    .when('/package/:packageName', {
	      controller:'PackageController',
	      templateUrl:'tpl/package.html'
	    })
	    .when('/package', {
	      controller:'PackageListController',
	      templateUrl:'tpl/package-list.html'
	    })
	    .otherwise({
	      redirectTo:'/'
    	});
})
})();
