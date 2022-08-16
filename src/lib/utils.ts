import { DateTime } from 'luxon';

export function getDateTime(date: string | Date) {
	if (typeof date === 'string') {
		const sFormat = 'yyyy-MM-dd HH:mm:ss';
		const msFormat = 'yyyy-MM-dd HH:mm:ss.SSS';
		const format = date.length === msFormat.length ? msFormat : sFormat;
		return DateTime.fromFormat(date, format, { zone: 'UTC' });
	}

	return DateTime.fromJSDate(date);
}

export function formatToUTCDate(date: string | Date, format = 'yyyy-MM-dd HH:mm:ss') {
	return getDateTime(date).toUTC().toFormat(format);
}

export function formatToLocalDate(date: string | Date, format = 'yyyy-MM-dd HH:mm:ss') {
	return getDateTime(date).toLocal().toFormat(format);
}

export function filtersFromSearchParams<Type extends { name: string }>(
	searchParams: URLSearchParams,
	filter: string,
	arr: Array<Type>
): { label: string; value: Type }[] {
	const valueMap = arr.reduce((map, element) => {
		map[element['name']] = element;
		return map;
	}, {} as { [key: string]: Type });

	const labels = searchParams.getAll(filter);

	return [
		...new Map(
			labels
				.map((label) => {
					return {
						label: label,
						value: valueMap[label]
					};
				})
				.map((option) => [option['label'], option])
		).values()
	];
}

export function post(endpoint: string, data: any) {
	return fetch(endpoint, {
		method: 'POST',
		credentials: 'include',
		body: JSON.stringify(data || {}),
		headers: {
			'Content-Type': 'application/json'
		}
	}).then((r) => r.json());
}
