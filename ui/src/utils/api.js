const GLOBAL_HEADERS = {};

export function setGlobalRequestHeader(header, value) {
  if (value === undefined) {
    delete GLOBAL_HEADERS[header];
  } else {
    GLOBAL_HEADERS[header] = value;
  }
}

export default async function api({
  url,
  method = 'GET',
  json = null,
  headers = {},
} = { url: '.' }) {
  const init = { method, headers: { ...headers } };

  for (const key of Object.keys(GLOBAL_HEADERS)) {
    if (!init.headers[key]) init.headers[key] = GLOBAL_HEADERS[key];
  }

  if (json !== null && method !== 'GET') {
    init.headers['Content-Type'] = 'application/json';
    init.body = JSON.stringify(json);
  }

  let response;
  try {
    response = await fetch(`/api/${url}`, init);
  } catch (e) {
    return { ok: false, status: 0, error: e };
  }

  let body = null;
  const text = await response.text().catch(() => null);
  if (text) {
    try { body = JSON.parse(text); } catch { body = text; }
  }

  return { ok: response.status >= 200 && response.status < 300, status: response.status, body };
}
