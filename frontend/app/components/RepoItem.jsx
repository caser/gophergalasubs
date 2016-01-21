var Actions = require('./Login.jsx');

var RepoItem = React.createClass({

  render: function(){
    return (<li>
            <h3>{this.props.repo.name}</h3>
            <p>
            {this.props.repo.description}
            </p>
            {this.props.repo.stargazers_count}
            <button onClick={this.props.vote}>Vote</button>
            </li>)
  }

});

module.exports = RepoItem;
