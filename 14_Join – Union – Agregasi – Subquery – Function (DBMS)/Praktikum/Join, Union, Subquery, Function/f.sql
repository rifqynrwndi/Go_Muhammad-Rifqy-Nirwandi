CREATE TRIGGER after_transaction_delete
AFTER DELETE ON transactions
FOR EACH ROW
BEGIN
   DELETE FROM transaction_details WHERE transaction_id = OLD.id;
END;
