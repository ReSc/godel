/**
 * context menu code derived from :
 *
 * ng-context-menu - v0.1.7 - An AngularJS directive to display a context menu
 * when a right-click event is triggered
 *
 * @author Ian Kennington Walter (http://ianvonwalter.com)

 */
(function () {
    angular
        .module('ng-context-menu', [])
        .factory('ContextMenuService', function () {
            return {
                element: null,
                menuElement: null
            };
        })
        .directive('contextMenu', [
            '$document',
            'ContextMenuService',
            function ($document, ContextMenuService) {
                return {
                    restrict: 'A',
                    scope: {
                        'callback': '&contextMenu',
                        'disabled': '&contextMenuDisabled'
                    },
                    link: function ($scope, $element, $attrs) {
                        var opened = false;

                        function open(event, menuElement) {
                            if (menuElement[0]) {
                                var doc = $document[0].documentElement;
                                var dl = (window.pageXOffset || doc.scrollLeft) - (doc.clientLeft || 0);
                                var dt = (window.pageYOffset || doc.scrollTop) - (doc.clientTop || 0);
                                var dw = doc.clientWidth + dl;
                                var dh = doc.clientHeight + dt;

                                var ew = menuElement[0].scrollWidth;
                                var eh = menuElement[0].scrollHeight;
                                var er = ew + event.pageX;
                                var eb = eh + event.pageY;
                                var el = Math.max(event.pageX - dl, 0);
                                var et = Math.max(event.pageY - dt, 0);

                                if (er > dw) {
                                    el = el - (er - dw);
                                }

                                if (eb > dh) {
                                    et = et - (eb - dh);
                                }

                                menuElement.css('top', et + 'px');
                                menuElement.css('left', el + 'px');
                                menuElement.css('visibility', 'visible');
                                menuElement.addClass('open');

                                opened = true;
                            } else {
                                alert("No contextmenu for " + $attrs.target);
                            }
                        }

                        function close(menuElement) {
                            menuElement.removeClass('open');
                            menuElement.css('visibility', 'hidden');
                            opened = false;
                        }

                        $element.bind('contextmenu', function (event) {
                            if (!$scope.disabled()) {
                                if (ContextMenuService.menuElement !== null) {
                                    close(ContextMenuService.menuElement);
                                }

                                ContextMenuService.menuElement = angular.element(
                                    document.getElementById($attrs.target)
                                );

                                ContextMenuService.element = event.target;
                                event.preventDefault();
                                event.stopPropagation();
                                $scope.$apply(function () {
                                    $scope.callback({ $event: event });
                                });

                                $scope.$apply(function () {
                                    open(event, ContextMenuService.menuElement);
                                });
                            }
                        });

                        function handleKeyUpEvent(event) {
                            if (!$scope.disabled() && opened && event.keyCode === 27) {
                                $scope.$apply(function () {
                                    close(ContextMenuService.menuElement);
                                });
                            }
                        }

                        function handleClickEvent(event) {
                            if (!$scope.disabled()
                                && opened
                                && (event.button !== 2 ||
                                    event.target !== ContextMenuService.element)) {
                                $scope.$apply(function () {
                                    close(ContextMenuService.menuElement);
                                });
                            }
                        }

                        $document.bind('keyup', handleKeyUpEvent);
                        // Firefox treats a right-click as a click and a contextmenu event
                        // while other browsers just treat it as a contextmenu event
                        $document.bind('click', handleClickEvent);
                        $document.bind('contextmenu', handleClickEvent);

                        $scope.$on('$destroy', function () {
                            $document.unbind('keyup', handleKeyUpEvent);
                            $document.unbind('click', handleClickEvent);
                            $document.unbind('contextmenu', handleClickEvent);
                        });
                    }
                };
            }
        ]);
})();