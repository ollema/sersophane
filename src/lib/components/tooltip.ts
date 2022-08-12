import tippy, { type Props } from 'tippy.js';

export function tooltip(node: HTMLElement, params: Partial<Props> = {}): SvelteActionReturnType {
	if (!tippy) return;

	// determine the title to show. prefer the custom content passed in first, then the HTML title attribute
	const custom = params.content;
	const title = node.title;
	const content = String(custom || title);

	node.title = '';
	const tip = tippy(node, { content, ...params });

	return {
		update: () => tip.setProps({ content, ...params }),
		destroy: () => tip.destroy()
	};
}
