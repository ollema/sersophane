import PocketBase from 'pocketbase';

declare global {
	namespace App {
		interface Locals {
			pocketbase: PocketBase;
		}
	}
}
