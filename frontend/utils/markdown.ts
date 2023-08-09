import DOMPurify from 'isomorphic-dompurify';
import * as marked from 'marked';

marked.use({
    mangle: false,
    headerIds: false
});

export default function renderMarkdown(text: string) {
    const md = marked.parse(text || '')
    return DOMPurify.isSupported ? DOMPurify.sanitize(md) : md;
}
