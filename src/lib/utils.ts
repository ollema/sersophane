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

export function formatToUTCDate(date: string | Date, format: string = 'yyyy-MM-dd HH:mm:ss') {
	return getDateTime(date).toUTC().toFormat(format);
}

export function formatToLocalDate(date: string | Date, format: string = 'yyyy-MM-dd HH:mm:ss') {
	return getDateTime(date).toLocal().toFormat(format);
}
