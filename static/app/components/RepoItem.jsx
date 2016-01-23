var RepoItem = React.createClass({

  render: function(){
    return (<li>
            <h3>{this.props.repo.name}</h3>
            <p>
            {this.props.repo.description}
            </p>
            {this.props.repo.stargazers_count} 
            <span className="glyphicon glyphicon-star" aria-hidden="true"></span>
            <button onClick={this.props.vote}>Vote</button>
            </li>)
  }

});

module.exports = RepoItem;
