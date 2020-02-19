# Copyright 2020 The Fedlearner Authors. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# coding: utf-8

import argparse
import logging

from data.data_block_set import DataBlockSet
from data.data_source_reader import DataSourceReader
from trainer_master import TrainerMaster


class FollowerTrainerMaster(TrainerMaster):
    def __init__(self, application_id, data_source_reader_):
        super(FollowerTrainerMaster, self).__init__(application_id)
        self._data_block_set = DataBlockSet()
        self._data_source_reader = data_source_reader_

    def _load_data(self):
        checkpoint = self._get_checkpoint()
        for data_block in self._data_source_reader.list_data_block():
            if data_block.block_id not in checkpoint:
                self._data_block_set.add(data_block)
        logging.debug("FollowerTrainerMaster: get all block %s",
                      self._data_block_set)

    def _alloc_data_block(self, block_id=None):
        logging.debug("FollowerTrainerMaster is getting block %s", block_id)
        if not block_id:
            raise Exception('follower tm need block_id to alloc.')
        return self._data_block_set.get(block_id)


if __name__ == '__main__':
    logging.getLogger().setLevel(logging.DEBUG)
    parser = argparse.ArgumentParser('leader trainer master cmd.')
    parser.add_argument('-p',
                        '--port',
                        type=int,
                        default=50002,
                        help='Listen port of follower trainer master')
    parser.add_argument('-app_id',
                        '--application_id',
                        required=True,
                        help='application_id')
    parser.add_argument('-data_path',
                        '--data_path',
                        required=True,
                        help='training example data path')
    parser.add_argument('-start_date',
                        '--start_date',
                        default=None,
                        help='training data start date')
    parser.add_argument('-end_date',
                        '--end_date',
                        default=None,
                        help='training data end date')
    FLAGS = parser.parse_args()
    data_source_reader = DataSourceReader(FLAGS.data_path, FLAGS.start_date,
                                          FLAGS.end_date)

    follower_tm = FollowerTrainerMaster(FLAGS.application_id,
                                        data_source_reader)
    follower_tm.run(listen_port=FLAGS.port)
