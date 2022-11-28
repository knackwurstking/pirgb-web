export default {
    /** @param {Event} ev */
    push(ev) {
        console.log("[router] push", ev)

        const path = ev.currentTarget.id
        self.history.pushState({ path: path }, `${path}`, `/${path}`)
    },

    enable() {
        console.log("[router] enable")

        self.addEventListener("popstate", (ev) => {
            console.log("[router] popstate:", ev, location)

            let path = ev.state.path || location.pathname.replace(/^\//, "")
            if (path) {
                // TODO: fire event || render content for path
            }
        })
    },
}
