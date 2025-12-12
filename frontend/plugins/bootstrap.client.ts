import bootstrap from 'bootstrap/dist/js/bootstrap.bundle'

export default defineNuxtPlugin(nuxtApp => {
    nuxtApp.provide('bootstrap', bootstrap)
    nuxtApp.hooks.hook('app:mounted', () => {
        const s = Object.assign(document.createElement('script'), {async: true, src: 'https://analytics.sakerhetssm.se/script.js'});
        s.dataset.websiteId = 'ec04f82f-8090-4b5a-b404-5a07099e1d78';
        document.body.append(s);
    });
})