package app

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/
func initRouter() map[string]string {

	routers := map[string]string{
		"index":    "/",
		"about-me": "/about-me/",
	}

	return routers
}
