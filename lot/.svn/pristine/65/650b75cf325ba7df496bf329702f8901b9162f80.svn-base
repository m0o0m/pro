<?php

namespace Applications\Common\Helper;

use MongoDB\Driver\BulkWrite;
use MongoDB\Driver\Exception\Exception;
use MongoDB\Driver\Manager;
use MongoDB\Driver\Query;
use MongoDB\Driver\WriteConcern;
use MongoDB\Driver\WriteResult;
use MongoDB\Driver\Command;
use MongoException;
use \Applications\Common\Config\Config as config;

class MongoDBManager {

    private $mongoManager;
    private $db;

    function __construct() {
        $mongoDBConfig = config::$mongo;
        $connectString = 'mongodb://';
        if ($mongoDBConfig['user'] && $mongoDBConfig['pass'])
            $connectString .= $mongoDBConfig['user'] . ':' . $mongoDBConfig['pass'] . '@';
        $connectString .= $mongoDBConfig['host'] . ':' . $mongoDBConfig['port'] . '/' . $mongoDBConfig['db'];
        $this->mongoManager = new Manager($connectString);
        $this->db = $mongoDBConfig['db'];
    }

    /**
     * @param string $collection
     * @param array $filter
     * @param array $options
     * @return array
     */
    public function executeQuery($collection, $filter = array(), $options = array()) {
        $query = new Query($filter, $options);
        return $this->mongoManager->executeQuery($this->db . '.' . $collection, $query)->toArray();
    }

    /**
     * @param string $collection
     * @param array $filter
     * @param array $options
     * @return int
     */
    public function count($collection, $filter = array(), $options = array()) {
        $command = new Command(["count" => $collection, "query" => $filter]);
        try {
            $result = $this->mongoManager->executeCommand($this->db, $command);
            $res = current($result->toArray());
            $count = $res->n;
        } catch (Exception $e) {
            echo $e->getMessage(), "\n";
        }
        return $count;
    }

    /**
     * @param string $collection
     * @param BulkWrite $bulkWrite
     * @return WriteResult
     */
    public function executeBulkWrite($collection, $bulkWrite) {
        return $this->mongoManager->executeBulkWrite($this->db . '.' . $collection, $bulkWrite);
    }

    /**
     * @param $doc
     * @param string $collection
     * @param bool $fetched
     * @return WriteResult
     */
    public function insertData($doc, $collection, $fetched = FALSE) {
        // do checking
        if (empty($doc) || $collection === NULL) {
            return false;
        }

        // save data information
        try {
            //$wc = new MongoDB\Driver\WriteConcern(MongoDB\Driver\WriteConcern::MAJORITY);

            $bulk = new BulkWrite();
            $insertedId = $bulk->insert($doc);
            $this->mongoManager->executeBulkWrite($this->db . '.' . $collection, $bulk);

            //throw new MongoException('insert data failed');

            if ($fetched) {
                return $insertedId;
            }
        } catch (Exception $e) {
            $this->throwError($e->getMessage());
        }
    }

    /**
     * Update records
     * @param $collection
     * @param $filter
     * @param $updated
     * @param $options
     * @return WriteResult
     */
    public function updateData($collection, $filter, $updated, $options = array()) {
        // do checking
        if ($collection === NULL || empty($updated) || empty($filter)) {
            $this->throwError('Updated data can not be empty!');
        }

        // do updating
        $timeout = 3000;
        $wc = new WriteConcern(WriteConcern::MAJORITY, $timeout);
        $bulk = new BulkWrite();
        $bulk->update($filter, $updated, $options);
        try {
            // execute
            return $this->mongoManager->executeBulkWrite("{$this->db}.$collection", $bulk, $wc);

            // throw new MongoException('find record failed');
        } catch (\MongoException $e) {
            $this->throwError($e->getMessage());
        }
    }

    /**
     * Delete record
     * @param $collection
     * @param $filter
     * @param $options
     * @return number of rows affected
     */
    public function deleteData($collection, $filter, $options = array()) {
        // do checking
        if ($collection === NULL) {
            $this->throwError('Inserted data can not be empty!');
        }

        if (!is_array($filter)) {
            $this->throwError('$filter format is invaild.');
        }

        try {
            // execute
            $bulk = new BulkWrite();
            $bulk->delete($filter, $options);
            $WriteResult = $this->mongoManager->executeBulkWrite("{$this->db}.$collection", $bulk);
            return $WriteResult->getDeletedCount();

            // throw new MongoException('delete record failed');
        } catch (MongoException $e) {
            $this->throwError($e->getMessage());
        }
    }

    /**
     * throw error message
     * @param string $errorInfo error message
     */
    private function throwError($errorInfo = '') {
        echo "<h3>Error：$errorInfo</h3>";
    }

}
