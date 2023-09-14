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
        fetch('http://localhost:8087/api/v1/volumes?getAll=true')
        .then(res => res.json())
        .then(
            (result) => {
                this.setState({
                    isLoaded: true,
                    items: result.Volumes
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
                            <th>DRIVER</th>
                            <th>VOLUME NAME</th>
                            <th>MOUNT POINT</th>
                            <th>SCOPE</th>
                            <th>CREATED</th>
                        </tr>
                    </thead>
                    <tbody>
                        {items.map(item => (
                            <tr key={item.Name}>
                                <td>{item.Driver}</td>
                                <td>{item.Name.substring(0, 12)} {item.Name >= 12 && '...'}</td>
                                <td>{item.Mountpoint}</td>
                                <td>{item.Scope}</td>
                                <td>{item.CreatedAt}</td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div>;
        }
    }
}

//ReactDOM.render(<App />, document.getElementById("volume"));
const root = ReactDOM.createRoot(document.getElementById("volume"));
root.render(<App />);