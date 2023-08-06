import * as DOMPurify from 'dompurify';
import * as marked from 'marked'

marked.use({
    mangle: false,
    headerIds: false
});

export default function renderMarkdown(text: string) {
    return DOMPurify.default.sanitize(marked.parse(text))
}
