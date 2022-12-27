CREATE TABLE grade_distribution (
    id SERIAL PRIMARY KEY,
    subject      VARCHAR(255) NOT NULL,
	sub_title     VARCHAR(255) NOT NULL,
	class        VARCHAR(255) NOT NULL,
	teacher      VARCHAR(255) NOT NULL,
	year         VARCHAR(255) NOT NULL,
	semester     VARCHAR(255) NOT NULL,
	faculty      VARCHAR(255) NOT NULL,
	student_count INTEGER NOT NULL,
	gpa          DECIMAL NOT NULL,

	a_count  INTEGER NOT NULL,
	ap_count INTEGER NOT NULL,
	am_count INTEGER NOT NULL,
	bp_count INTEGER NOT NULL,
	b_count  INTEGER NOT NULL,
	bm_count INTEGER NOT NULL,
	cp_count INTEGER NOT NULL,
	c_count  INTEGER NOT NULL,
	d_count  INTEGER NOT NULL,
	dm_count INTEGER NOT NULL,
	f_count  INTEGER NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE FUNCTION set_update_time() RETURNS TRIGGER AS '
  begin
    new.updated_at := ''now'';
    return new;
  end;
' LANGUAGE plpgsql;

CREATE TRIGGER update_tri BEFORE UPDATE ON grade_distribution FOR EACH ROW EXECUTE PROCEDURE set_update_time();