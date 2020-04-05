package mock

const WeatherStackResp = `{"request":{"type":"City","query":"Melbourne, Australia","language":"en","unit":"m"},"location":{"name":"Melbourne","country":"Australia","region":"Victoria","lat":"-37.817","lon":"144.967","timezone_id":"Australia\/Melbourne","localtime":"2020-04-06 01:00","localtime_epoch":1586134800,"utc_offset":"10.0"},"current":{"observation_time":"03:00 PM","temperature":11,"weather_code":116,"weather_icons":["https:\/\/assets.weatherstack.com\/images\/wsymbols01_png_64\/wsymbol_0004_black_low_cloud.png"],"weather_descriptions":["Partly cloudy"],"wind_speed":13,"wind_degree":280,"wind_dir":"W","pressure":1016,"precip":0.4,"humidity":87,"cloudcover":75,"feelslike":9,"uv_index":1,"visibility":10,"is_day":"no"}}`

const OpenWeatherResp = `{"coord":{"lon":-80.61,"lat":28.08},"weather":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04d"}],"base":"stations","main":{"temp":294.21,"feels_like":293.69,"temp_min":293.15,"temp_max":295.15,"pressure":1017,"humidity":73},"visibility":16093,"wind":{"speed":3.6,"deg":310},"clouds":{"all":90},"dt":1586099337,"sys":{"type":1,"id":4922,"country":"US","sunrise":1586084844,"sunset":1586130128},"timezone":-14400,"id":4163971,"name":"Melbourne","cod":200}`