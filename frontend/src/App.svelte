<script>
    import {
        Router,
        Route,
        links,
        navigate as svelteNavigate,
    } from "svelte-routing";

    import Home from "./pages/Home.svelte";
    import Sections from "./pages/Sections.svelte";
    import Groups from "./pages/Groups.svelte";
    import Settings from "./pages/Settings.svelte";

    import StatusLED from "./lib/StatusLED.svelte";

    let currentUrl = location.pathname;
    let open = false;

    /** @param {string} url */
    function navigate(url) {
        svelteNavigate((currentUrl = url));
    }
</script>

<main use:links>
    <Router>
        <nav>
            <div class="left">
                <StatusLED active={open} />
            </div>

            <div class="right">
                <button
                    class="link"
                    class:active={currentUrl === "/sections"}
                    on:click={() => navigate("/sections")}>Home</button
                >
                <button
                    class="link"
                    class:active={currentUrl === "/groups"}
                    on:click={() => navigate("/groups")}>Groups</button
                >
                <button
                    class="link"
                    class:active={currentUrl === "/settings"}
                    on:click={() => navigate("/settings")}>Settings</button
                >
            </div>
        </nav>
        <div class="container">
            <Route path="/sections">
                <Sections bind:open />
            </Route>
            <Route path="/groups" component={Groups} />
            <Route path="/settings" component={Settings} />
            <Route path="/"><Home {navigate} /></Route>
        </div>
    </Router>
</main>

<style>
    :global(body) {
        scroll-behavior: smooth;
    }

    nav {
        height: calc(1rem + 25px);
        display: flex;
        place-items: center;
        justify-content: flex-end;
        border-bottom: 1px solid var(--color-border);
    }

    nav .left {
        width: fit-content;
    }

    nav .right {
        width: 100%;
        display: flex;
        place-items: center;
        justify-content: flex-end;
    }

    nav button.link {
        background: transparent;
        border: none;
        outline: none;
        margin: 4px 8px;
        padding: 6px 16px;
        color: var(--color-link);
        transition: color 0.25s linear;
    }

    nav button.link.active {
        color: var(--color-link--active);
    }
</style>
