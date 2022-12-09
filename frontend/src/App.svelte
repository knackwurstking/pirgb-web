<script>
  import { Router, Route, links, navigate as _navigate } from "svelte-routing";

  import Home from "./pages/Home.svelte";
  import Sections from "./pages/Sections.svelte";
  import Groups from "./pages/Groups.svelte";
  import Settings from "./pages/Settings.svelte";

  let currentUrl = location.pathname;

  /** @param {string} url */
  function navigate(url) {
    _navigate((currentUrl = url));
  }
</script>

<main use:links>
  <Router>
    <nav>
      <button
        class="link"
        class:active={currentUrl === "/sections"}
        on:click={() => navigate("/sections")}
      >Home</button>
      <button
        class="link"
        class:active={currentUrl === "/groups"}
        on:click={() => navigate("/groups")}
      >Groups</button>
      <button
        class="link"
        class:active={currentUrl === "/settings"}
        on:click={() => navigate("/settings")}
      >Settings</button>
    </nav>
    <div>
      <Route path="/sections" component={Sections} />
      <Route path="/groups" component={Groups} />
      <Route path="/settings" component={Settings} />
      <Route path="/" component={Home} />
    </div>
  </Router>
</main>

<style>
  nav {
    height: calc(1rem + 25px);
    display: flex;
    place-items: center;
    justify-content: flex-end;
    border-bottom: 1px solid currentColor;
  }

  nav button.link {
    background: transparent;
    border: none;
    outline: none;
    margin: 4px 16px;
    color: var(--color-link);
  }

  nav button.link.active {
    color: var(--color-link--active);
  }
</style>
