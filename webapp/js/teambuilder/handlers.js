window.TeamActionHandler = {
  initialize: function (teams) {
    this._teams = teams;
    this._teams.addBunchFromData(this._loadTeamData())
  },

  _loadTeamData: function () {
    return (JSON.parse(window.localStorage.getItem("mmo_teams"))) || [];
  },

  create: function () {
    this._teams.addStore(new Team({
      id: this._teams.stores.length + 1,
      name: "unnamed"
    }));
  }
}
