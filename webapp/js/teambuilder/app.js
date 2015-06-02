window.TeamBuilder = React.createClass({
  mixins: [Fluxo.WatchComponent],

  listenProps: ["teams"],

  createNewTeam: function () {
    Fluxo.callAction("Teams", "create");
  },

  render: function () {
    return (
      <div>
        <TeamList teams={this.props.teams} />
        <button onClick={this.createNewTeam} className="btn btn-lg btn-primary">
          <i className="glyphicon glyphicon-plus"></i>
          New team
        </button>
      </div>
    )
  }
});
