class Menu extends React.Component {

    constructor(props) {
        super(props);
        this.state = {routeState: '/containers'};
        // This binding is necessary to make `this` work in the callback
        this.handleClick = this.handleClick.bind(this);
    }

    handleClick(menuURL) {
        this.setState({routeState: menuURL});
    }

    render() {
        const menus = [
            {"menuName": "Containers", "menuURL": "/containers"},
            {"menuName": "Volumes", "menuURL": "/volumes"},
            {"menuName": "Images", "menuURL": "/images"},
            {"menuName": "Help", "menuURL": "/help"}
        ]
    
        return <div className="col-md-3">
            <div className="list-group">
                {menus.map((menu) =>
                    <a href={menu.menuURL} onClick={this.handleClick.bind(this, menu.menuURL)} className="list-group-item" key={menu.menuName}>{menu.menuName}</a>
                )}
            </div>
        </div>;
    }
}