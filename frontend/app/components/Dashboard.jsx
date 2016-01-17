var Dashboard = React.createClass({
  render: function() {
    var repos = this.props.state.repos.map(function(repo){
      return (<li key={repo.id}>
              <h3>{repo.name}</h3>
              <p>
              {repo.description}
              </p>
              {repo.stargazers_count}
              <button>Vote</button>
              </li>)
    })
    return (
      <div>
        <div id="passport">
          {this.props.state.user.login}
           <img src={this.props.state.user.avatar_url} width="50px"/>
        </div>
        <h2>Dashboard</h2>
        <ul>
          {repos}
        </ul>
        <div>
        <h2>Your top 5</h2>
        <ol>
        <li>
        item
        </li>
        <li>
        item
        </li>
        <li>
        item
        </li>
        <li>
        item
        </li>
        <li>
        item
        </li>
        </ol>
        </div>
      </div>
    );
  }
});

module.exports = Dashboard;
