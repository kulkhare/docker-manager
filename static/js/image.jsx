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

    home() {
        alert("clicked..")
    }
}

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
        const apiEndPoint = 'api/v1';
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

class ContentWindow extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            error: null,
            isLoaded: false,
            items: []
        };
    }

    componentDidMount() {
        fetch('http://localhost:8087/api/v1/images?getAll=true')
        .then(res => res.json())
        .then(
            (result) => {
                this.setState({
                    isLoaded: true,
                    items: result
                });
            },
            // Note: it's important to handle errors here
            // instead of a catch() block so that we don't swallow
            // exceptions from actual bugs in components.
            (error) => {
                this.setState({
                    isLoaded: true,
                    error
                });
            }
        )
    }

    render() {
        const { error, isLoaded, items } = this.state;
        if (error) {
            return <div>Error: {error.message}</div>;
        } else if (!isLoaded) {
            return <div>Loading...</div>;
        } else {
            return <div className="col-md-9">
                <table className="table">
                    <thead>
                        <tr>
                            <th>REPOSITORY</th>
                            <th>IMAGE ID</th>
                            <th>CREATED</th>
                            <th>SIZE</th>
                        </tr>
                    </thead>
                    <tbody>
                        {items.map(item => (
                            <tr key={item.Id}>
                                <td>{item.RepoTags}</td>
                                <td>{item.Id.substring(0, 12)} {item.Id.length >= 12 && '...'}</td>
                                <td>{item.Created} </td>
                                <td>{item.Size} </td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div>;
        }
    }
}

//ReactDOM.render(<App />, document.getElementById("image"));
const root = ReactDOM.createRoot(document.getElementById("image"));
root.render(<App />);