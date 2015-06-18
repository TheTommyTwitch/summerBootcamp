
"use strict";

var artistControllers = angular.module('artistControllers', []);

artistControllers.controller('ListController', ['$scope', 'GetData', function($scope, GetData) {
    $scope.artists = GetData;
    $scope.artistOrder = 'name';
}]);


artistControllers.controller('DetailsController', 
	['$scope', 'GetData','$routeParams',
	function($scope, GetData, $routeParams) {
		$scope.artists = GetData;
    $scope.whichItem = $routeParams.itemId;
}]);

//artistControllers.controller('DetailsController', ['$scope', '$http', '$routeParams', function($scope, $http, $routeParams) {
//  $http.get('js/data.json').success(function(data) {
//    $scope.artists = data;
//    $scope.whichItem = $routeParams.itemId;
//  });
//}]);