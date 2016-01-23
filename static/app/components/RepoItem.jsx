var RepoItem = React.createClass({

  render: function(){
    return (<li className="repo">
            <div className="details">
              <h3 className="name">
                {this.props.repo.name}
                <span className="stargazers" title={this.props.repo.stargazers_count + " Stargazers"}>
                  <span className="glyphicon glyphicon-star" aria-hidden="true"></span>
                  {this.props.repo.stargazers_count} 
                </span>
              </h3>
              <p className="description">{this.props.repo.description || "This submission has no description."}</p>
            </div>
            <button className="btn btn-green" onClick={this.props.vote}>
              Vote <span className="glyphicon glyphicon-arrow-up" aria-hidden="true"></span>
            </button>
            </li>)
  }

});

module.exports = RepoItem;
