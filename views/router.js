export default {
    /** @param {PopStateEvent} ev */
    popstate(ev) {
        console.log("[router] popstate", ev, this)
        // TODO: get the id from state (or url/path); render content
    },

    /** @param {Event} ev */
    push(ev) {
        console.log("[router] push", ev)

        self.history.pushState(
            { id: ev.currentTarget.id },
            `${ev.currentTarget.id}`,
            `/${ev.currentTarget.id}`,
        )
    },

    enable() {
        console.log("[router] enable")
        self.addEventListener("popstate", this.popstate)
    },

    disable() {
        console.log("[router] disable")
        self.removeEventListener("popstate", this.popstate)
    },
}
