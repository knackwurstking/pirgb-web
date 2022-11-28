/** @type {(ev: PopStateEvent) => void} */
const onpopstate = (ev) => {
    console.log("[router] popstate", ev)
}

/** @param {Event} ev */
function push(ev) {
    console.log("[router] push", ev)
    history.pushState(
        { id: ev.currentTarget.id },
        `${ev.currentTarget.id}`,
        `/${ev.currentTarget.id}`,
    )
}

function enable() {
    window.addEventListener("popstate", onpopstate)
}

function disable() {
    window.removeEventListener("popstate", onpopstate)
}

export default {
    push,
    enable,
    disable,
}
