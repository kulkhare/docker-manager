class App extends React.Component {
    render() {
        const appName = 'Docker Manager';
        return (
            <div className="container">
                <br />
                <span className="pull-right">
                    <a onClick={this.home}>Home</a>
                </span>
                <h2>{ appName }</h2>
                <p>Docker Manger is an easy-to-install web application It provides a simple interface that enables you to manage your containers, applications, and images directly from your machine without having to use cli.</p>
                <div className="row">
                    <Menu />
                    <ContentWindow />
                </div>
            </div>
        );
    }
}

class Menu extends React.Component {
    render() {
        const apiEndPoint = 'api/v1';
        const menus = [
            {"menuName": "Containers", "menuURL": "/containers"},
            {"menuName": "Volumes", "menuURL": "/volumes"},
            {"menuName": "Images", "menuURL": "/images"},
            {"menuName": "Help", "menuURL": "/help"}
        ]

        return <div className="col-md-3">
            <ListItem menus={menus} />
        </div>;
        
        function ListItem(props) {
            const menus = props.menus;
            return (
                <div className="list-group">
                    {menus.map((menu) =>
                        <a href={apiEndPoint + menu.menuURL} className="list-group-item" key={menu.menuName}>{menu.menuName}</a>
                    )}
                </div>
            );
        }
    }

}

class ContentWindow extends React.Component {
    render() {
        return <div className="col-md-9">

        </div>;
    }
}

//ReactDOM.render(<App />, document.getElementById("app"));
const root = ReactDOM.createRoot(document.getElementById("app"));
root.render(<App />);