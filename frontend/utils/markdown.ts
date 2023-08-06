import * as DOMPurify from 'dompurify';
import * as marked from 'marked'

export default function renderMarkdown(text: string) {
    return DOMPurify.default.sanitize(marked.parse(text))
}
