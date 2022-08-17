import PocketBase, { User } from 'pocketbase';

declare global {
	namespace App {
		interface Locals {
			pocketbase: PocketBase;
			user?: User;
		}
	}
}
