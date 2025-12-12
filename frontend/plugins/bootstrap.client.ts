import bootstrap from 'bootstrap/dist/js/bootstrap.bundle'

export default defineNuxtPlugin(nuxtApp => {
    nuxtApp.provide('bootstrap', bootstrap)
    nuxtApp.hooks.hook('page:finish', () => {
        const c = new AbortController();
        const t = setTimeout(() => c.abort(), 2500);
        fetch('https://analytics.sakerhetssm.se/script.js', {signal: c.signal})
            .then(r => r.text())
            .then(code => {
                const s = Object.assign(document.createElement('script'), {textContent: code});
                s.dataset.websiteId = 'ec04f82f-8090-4b5a-b404-5a07099e1d78';
                document.body.append(s);
            })
            .catch(() => {})
            .finally(() => clearTimeout(t));
    });
})