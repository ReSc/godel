//typescipt



module TreeView {
	// the zero id is the root of all trees
	// no application node should ever have the zero id.
	export class Node {

		private _isSelected: boolean;

		get isSelected(): boolean {
			return this._isSelected;
		}

		set isSelected(v: boolean) {
			if (v != this._isSelected) {
				this._isSelected = v;
				this.context.onSelectedChanged(this);
			}
		}

		private _isVisible: boolean;

		get isVisible(): boolean {
			return this._isVisible;
		}

		set isVisible(v: boolean) {
			if (v != this._isVisible) {
				this._isVisible = v;
				this.context.onVisibleChanged(this);
			}
		}

		get parent(): Node {
		 return this.context.getNode(this.parentId)
	 }

		get siblings(): Node[] {
		 return this.context.getChildren(this.parent)
	 }

		get children(): Node[] {
		 return this.context.getChildren(this)
	 }

		constructor(
			public context: Context,
			public id: any,
			public parentId: any,
			public label: string,
			public icon: string,
			public data: any
			) {
			this._isSelected = false;
			this._isVisible = true;
		}

		select() {
			this.isSelected = true;
		}

		hide(): void {
			this.isVisible = false;
		}

		show(): void {
			this.isVisible = true;
		}

		refresh(): void {
			this.context.onRefresh(this);
		}

		refreshChildren(): void {
			this.context.onRefreshChildren(this);
		}
	}

	export class Context {

		private _selectedNodes: {};

		constructor(
			public name: string
			) {
			this._selectedNodes = {};
		}

		onSelectedChanged(node: Node): void {
			if (this._selectedNodes[node.id]) {
				if (!node.isSelected) {
					delete this._selectedNodes[node.id];
				}
			} else {
				if (node.isSelected) {
					this._selectedNodes[node.id] = node;
				}
			}
		}

		onVisibleChanged(node: Node): void {

		}

		onRefresh(node: Node): void {

		}

		onRefreshChildren(node: Node): void {

		}

		getSelectedNodes(): Node[] {
			var nodes = new Array<Node>();
			for (var key in this._selectedNodes) {
				if (this._selectedNodes.hasOwnProperty(key)) {
					nodes.push(this._selectedNodes[key]);
				}
			}
			return nodes;
		}

		getNode(id: any): Node {
		return null
	 }

		getChildren(node: Node): Node[] {
		return []
	 }

		createNode(id: any, parentId: any, label: string, icon: string, data: any): Node {
			return new Node(this, id, parentId, label, icon, data);
		}
	}
}

var ctx1 = new TreeView.Context("ctx 1");
var n1 = ctx1.createNode(1, 0, "node 1", "", "");
var n2 = ctx1.createNode(2, 0, "node 2", "", "");
var n3 = ctx1.createNode(3, 0, "node 3", "", "");
n1.select();
n2.select();
n3.select();
n1.isSelected = false;
n2.isSelected = false;
n3.isSelected = false;


