const iframe = document.createElement('iframe');
iframe.style.display = 'none';
document.body.appendChild(iframe);

let loaded = false;

iframe.onload = () => {
  loaded = true;
  
  // Try to access the frame's content
  try {
    const content = iframe.contentDocument || iframe.contentWindow.document;
    if (content && content.body) {
      console.log('Site is frameable - can access content');
    }
  } catch (e) {
    console.log('Site is NOT frameable - blocked from accessing content');
  }
  
  iframe.remove();
};

setTimeout(() => {
  if (!loaded) {
    console.log('Site is NOT frameable - onload never fired');
    iframe.remove();
  }
}, 3000);

iframe.src = 'https://google.com';
