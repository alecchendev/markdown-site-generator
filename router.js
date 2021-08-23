const routes = [{path: `/example`,template: `<h2 class="header">This is a different page</h2>

<ul>
<li>Bullet</li>
<li>Points</li>
</ul>

<p class="image"><img src="static/image.jpg" alt="Image of butter dog" /></p>
`,},{path: `/index`,template: `<h1>Hello, welcome to the page!</h1>

<p>Plaintext sample. Lorem ipsum.</p>
`,},];class Router {

	constructor(routes) {
		this.routes = routes;
		this._loadInitialRoute();
	}

	loadRoute(...urlSegments) {
		// Attempt to match the URL to a route.
		const matchedRoute = this._matchUrlToRoute(urlSegments);

		// Push a history entry with the new URL.
		// We pass an empty object and an empty string as the historyState
		// and title arguments, but their values do not really matter here.
		const url = `/${urlSegments.join('/')}`;
		history.pushState({}, '', url);

		// Append the template of the matched route to the DOM, 
		// inside the element with attribute data-router-outlet.
		const routerOutletElement = document.getElementById("content");
		routerOutletElement.innerHTML = matchedRoute.template;
	}

	_matchUrlToRoute(urlSegments) {
		// Try and match the URL to a route.
		const matchedRoute = this.routes.find(route => {

			// We assume that the route path always starts with a slash, and so 
			// the first item in the segments array  will always be an empty
			// string. Slice the array at index 1 to ignore this empty string.
			const routePathSegments = route.path.split('/').slice(1);

			// If there are different numbers of segments, then the route 
			// does not match the URL.
			if (routePathSegments.length !== urlSegments.length) {
				return false;
			}

			// If each segment in the url matches the corresponding route path, 
			// then the route is matched.
			return routePathSegments
				.every((routePathSegment, i) => routePathSegment === urlSegments[i]);
		});
		return matchedRoute;
	}

	_loadInitialRoute() {
		// Figure out the path segments for the route which should load initially.
		const pathnameSplit = window.location.pathname.split('/');
		const pathSegments = pathnameSplit.length > 1 ? pathnameSplit.slice(1) : '';

		// Load the initial route.
		this.loadRoute(...pathSegments );
	}
}

export const router = new Router(routes);