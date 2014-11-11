(function(){
	var app= angular.module("godel",['ngResource','ngMaterial','ng-context-menu'])

	app.controller("TreeController",
		['$scope', '$resource',
		function($scope, $resource){

			var TreeNode = $resource("/api/v1/node/:id",{id:'@id'})
			$scope.nodes = TreeNode.query({parentId:0})

			$scope.toggleExpanded= function(e, node){
				e.stopPropagation();
				if(node.isExpanded){
					node.isExpanded=false;
				} else {
					node.isExpanded = true;
					if (node.nodes){
						return;
					}
					var children = TreeNode.query(
						{ parentId: node.id },
						function(){
						for (var i =0; i<children.length;i++){
							children[i].menuitems = [{
								id: children[i].id +"_create_child",
								name: "Create New Node"
							},{
								id: children[i].id +"_delete_node",
								name: "Remove Node"
							}];
						}
						node.nodes = children;
					});
				}
			};


		}]);

	app.controller("CreateNodeController",
		['$scope', '$resource',
		function($scope, $resource){
			var TreeNode = $resource("/api/v1/node/:id",{id:'@id'})
			$scope.placeholder="node name"
			$scope.name= "New Node"
			$scope.submit = function(){
				var newNode = new TreeNode({name:$scope.name})
				newNode.$save(
					function(value, responseHeaders){
						console.log(responseHeaders);
					},
					function(httpResponse){
						console.log(httpResponse);
					}
				)
			}
		}
	]);
	app.factory('selectionServiceFactory', function(){
		var	selectedItems = [];

		return {

			primarySelection: function(){
				if (selectedItems.length>0){
					return selectedItems[0]
				}
				return null;
			},

			selection: function(){
				return selectedItems.slice(0)
			},

			addSelection: function(item){
				for (var i = 0; i < selectedItems.length; i++){
					if selectedItems[i] === item {
						return;
				}
				selectedItems.push(item);
			},

			setPrimarySelection: function(item){
				if (selectedItems.length > 0 && selectedItems[0] === item){
					return;
				}
				for (var i = 1; i < selectedItems.length; i++){
					if selectedItems[i] === item {
						selectedItems.splice(i,1);
						break;
					}
				}
				selectedItems.unshift(item)
			},

			removeSelection: function(item){
				for (var i = 0; i < selectedItems.length; i++){
					if selectedItems[i] === item {
						selectedItems.splice(i, 1);
						return;
					}
				}
			},

			clear: function(){
				selectedItems=[];
			},
		};
	});
})();
