# Notes

## Queries

## PreBook verify
```shell
bioletus_dev=# select reservation_id, reserved_by_id, reserved_at, status from tickets where status = 'reserved';
 reservation_id |            reserved_by_id            |        reserved_at         |  status
 ----------------+--------------------------------------+----------------------------+----------
  267bd14dbf78   | dbde4763-7e9d-4d43-94af-64230d398d44 | 2020-01-17 01:59:15.116746 | reserved
  267bd14dbf78   | dbde4763-7e9d-4d43-94af-64230d398d44 | 2020-01-17 01:59:15.116746 | reserved
  267bd14dbf78   | dbde4763-7e9d-4d43-94af-64230d398d44 | 2020-01-17 01:59:15.116746 | reserved
  267bd14dbf78   | dbde4763-7e9d-4d43-94af-64230d398d44 | 2020-01-17 01:59:15.116746 | reserved
(4 rows)
```




