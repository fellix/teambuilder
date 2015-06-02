window.Ability = Fluxo.Store.extend({})

window.Item = Fluxo.Store.extend({})

window.Pokemon = Fluxo.Store.extend({})

window.Move = Fluxo.Store.extend({})

window.TeamMember = Fluxo.Store.extend({})

window.Team = Fluxo.Store.extend({})

window.Teams = Fluxo.CollectionStore.extend({
  store: window.Team
})
