# coding: utf-8

"""
    Kubeflow Pipelines API

    This file contains REST API specification for Kubeflow Pipelines. The file is autogenerated from the swagger definition.

    Contact: kubeflow-pipelines@google.com
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest

import kfp_server_api
from kfp_server_api.api.pipeline_service_api import PipelineServiceApi  # noqa: E501
from kfp_server_api.rest import ApiException


class TestPipelineServiceApi(unittest.TestCase):
    """PipelineServiceApi unit test stubs"""

    def setUp(self):
        self.api = kfp_server_api.api.pipeline_service_api.PipelineServiceApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_create_pipeline(self):
        """Test case for create_pipeline

        Creates a pipeline.  # noqa: E501
        """
        pass

    def test_create_pipeline_and_version(self):
        """Test case for create_pipeline_and_version

        Creates a new pipeline and a new pipeline version in a single transaction.  # noqa: E501
        """
        pass

    def test_create_pipeline_version(self):
        """Test case for create_pipeline_version

        Adds a pipeline version to the specified pipeline ID.  # noqa: E501
        """
        pass

    def test_delete_pipeline(self):
        """Test case for delete_pipeline

        Deletes an empty pipeline by ID. Returns error if the pipeline has pipeline versions.  # noqa: E501
        """
        pass

    def test_delete_pipeline_version(self):
        """Test case for delete_pipeline_version

        Deletes a specific pipeline version by pipeline version ID and pipeline ID.  # noqa: E501
        """
        pass

    def test_get_pipeline(self):
        """Test case for get_pipeline

        Finds a specific pipeline by ID.  # noqa: E501
        """
        pass

    def test_get_pipeline_by_name(self):
        """Test case for get_pipeline_by_name

        Finds a specific pipeline by name and namespace.  # noqa: E501
        """
        pass

    def test_get_pipeline_version(self):
        """Test case for get_pipeline_version

        Gets a pipeline version by pipeline version ID and pipeline ID.  # noqa: E501
        """
        pass

    def test_list_pipeline_versions(self):
        """Test case for list_pipeline_versions

        Lists all pipeline versions of a given pipeline ID.  # noqa: E501
        """
        pass

    def test_list_pipelines(self):
        """Test case for list_pipelines

        Finds all pipelines within a namespace.  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
