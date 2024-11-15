package repository

const (
	InsertQuery                        = `INSERT into link_info(link_id, full_link, shortner_link, visits, last_visit_time, created_at)  values ($1,$2, $3, $4, $5, $6)  returning id `
	GetQuery                           = `SELECT link_id, full_link, shortener_link, visits, last_visit_time,  created_at FROM link_info`
	UpdateVisitorsByShortenerLinkQuery = `UPDATE link_info SET visits = visits + 1, xf = $1 WHERE shortener_link = $2`
	DeleteByShortenerLinkQueryQuery    = `DELETE FROM link_info WHERE shortener_link = $1`
	GetByShortenerLinkQuery            = `SELECT link_id, full_link, shortener_link, visits, last_visit_time, created_at  FROM link_info  WHERE shortener_link = $1`
	DeleteExpiredShortenerLinkQuery    = `DELETE FROM link_info  WHERE created_at < NOW() - INTERVAL 30 DAY`
)
