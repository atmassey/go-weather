document.getElementById('search-form').addEventListener('submit', function(event) {
    event.preventDefault();
    var city = document.getElementById('city-input').value;
    fetch('/weather?city=' + city)
        .then(response => response.json())
        .then(data => {
            displayWeather(data);
        })
        .catch(error => {
            console.error('Error fetching weather:', error);
        });
});

function displayWeather(weatherData) {
    var weatherInfoDiv = document.getElementById('weather-info');
    weatherInfoDiv.innerHTML = `
        <h2>${weatherData.city}</h2>
        <div class="forecast">
            ${weatherData.forecast.map(day => `
                <div class="day-forecast">
                    <h3>${day.date}</h3>
                    <p>Temperature: ${day.temperature}°C</p>
                    <p>Weather: ${day.weather}</p>
                    <p>Wind Speed: ${day.windSpeed} m/s</p>
                </div>
            `).join('')}
        </div>
    `;
}
