(function (window) {
  var teams = new Teams;
  
  Fluxo.registerActionHandler("Teams", window.TeamActionHandler, teams);

  React.render(
    <TeamBuilder teams={teams} />,
    window.document.getElementById("app")
  )
})(window);
