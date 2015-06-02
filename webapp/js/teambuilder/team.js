window.TeamItem = React.createClass({
  mixins: [Fluxo.WatchComponent],

  listenProps: ["team"],

  render: function () {
    return (
      <li role="presentation"><a href='#'>{this.state.team.name}</a></li>
    )
  }
})

window.TeamList = React.createClass({
  mixins: [Fluxo.WatchComponent],

  listenProps: ["teams"],

  render: function () {
    var teams = this.state.teams;

    if (_.isEmpty(teams.stores)) {
      return (
        <p>No team was found</p>
      )
    } else {
      return (
        <ul className='nav nav-tabs'>
          {_.map(teams, function (team) {
            return <TeamItem key={team.id} team={team} />
          })}
        </ul>
      )
    }
  }
})
