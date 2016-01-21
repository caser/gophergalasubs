var RepoItem = require('./RepoItem.jsx')
var Actions = require('../actions.js')

var Dashboard = React.createClass({

  vote: function(id) {
    Actions.submitVote(this.props.dispatch, id)
  },

  render: function() {
    var vote = this.vote;
    var repos = this.props.state.repos.map(function(repo){
      return <RepoItem key={repo.id} repo={repo} vote={function(){vote(repo.id)}} />
    })

    return (
      <div>
        <div id="passport" className="row">
          <div id="passport" className="col-md-8">
            <h1>Gopher Gala Voting</h1>
          </div>
          <div id="passport" className="col-md-4">
            {this.props.state.user.login}
             <img src={this.props.state.user.avatar_url} width="50px"/>
          </div>
        </div>
        <div className="row">
          <div className="col-md-8">
            <h2>Submissions</h2>
            <ul>
              {repos}
            </ul>
          </div>
          <div className="col-md-4">
            <h2>Your top 5</h2>
            <ol>
              <li>item</li>
              <li>item</li>
              <li>item</li>
              <li>item</li>
              <li>item</li>
            </ol>
          </div>
        </div>
      </div>
    );
  }
});

module.exports = Dashboard;
