// document.addEventListener('DOMContentLoaded', function() {
//     fetch('http://localhost:8000/api/ad')
//         .then(response => response.json())
//         .then(data => {
//             document.getElementById('ad-title').textContent = data.title;
//             const adImage = document.getElementById('ad-image');
//             adImage.src = data.image_url;

//             // ثبت رویداد نمایش
//             fetch(data.impression_event, {
//                 method: 'POST',
//                 headers: {
//                     'Content-Type': 'application/json'
//                 },
//                 body: JSON.stringify({
//                     title: data.title,
//                     url: data.image_url
//                 })
//             })
//             .catch(error => console.error('Error sending impression event:', error));

//             // ثبت رویداد کلیک
//             adImage.addEventListener('click', function() {
//                 fetch(data.click_event, {
//                     method: 'POST',
//                     headers: {
//                         'Content-Type': 'application/json'
//                     },
//                     body: JSON.stringify({
//                         title: data.title,
//                         url: data.image_url
//                     })
//                 })
//                 .catch(error => console.error('Error sending click event:', error));
//             });
//         })
//         .catch(error => {
//             console.error('Error fetching ad:', error);
//             document.getElementById('ad-title').textContent = 'Failed to load ad';
//         });
// });

async function fetchAd() {
    try {
        const response = await fetch('http://localhost:9005/api/ad'); // Ensure this URL matches the server
        const ad = await response.json();
        displayAd(ad);
        registerImpressionEvent(ad.impression_event);
    } catch (error) {
        console.error('Error fetching ad:', error);
    }
}
