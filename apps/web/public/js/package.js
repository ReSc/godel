(function(){
	angular.
	module("godel").
	factory("Packages", function($resource){
		return $resource("/api/v1/model/package/:id")
	}).
	controller("PackageListController",
	function($scope, $location, Packages) {
		$scope.packages = Packages.query()
	}).
	controller("PackageController",
	function($scope, $location, $routeParams, Packages) {
		var name = $routeParams.packageName
		$scope.name = name
		$scope.package = Packages.get({'id':name})
	})
})();

