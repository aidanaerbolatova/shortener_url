package repository

const (
	InsertQuery                        = `INSERT into link_info(full_link, shortener_link, visits, last_visit_at, created_at)  values ($1, $2, $3, $4, $5)  returning id `
	GetQuery                           = `SELECT id, full_link, shortener_link, visits, last_visit_at,  created_at FROM link_info`
	UpdateVisitorsByShortenerLinkQuery = `UPDATE link_info SET visits = visits + 1 WHERE shortener_link = $1`
	DeleteByShortenerLinkQueryQuery    = `DELETE FROM link_info WHERE shortener_link = $1`
	GetByShortenerLinkQuery            = `SELECT id, full_link, shortener_link, visits, last_visit_at, created_at  FROM link_info  WHERE shortener_link = $1`
	DeleteExpiredShortenerLinkQuery    = `DELETE FROM link_info WHERE created_at < NOW() - INTERVAL '30 days'`
)
