const zeroPad = (num: number, places: number) => String(num).padStart(places, '0');

export function formatDate(date: string | Date) {
	const d = new Date(date);
	return `${zeroPad(d.getDate(), 2)}/${zeroPad(d.getMonth() + 1, 2)}`;
}
