CREATE TRIGGER after_transaction_detail_delete
AFTER DELETE ON transaction_details
FOR EACH ROW
BEGIN
   UPDATE transactions
   SET total_qty = total_qty - OLD.qty
   WHERE id = OLD.transaction_id;
END;
