-- Initial database schema for VibeStorm Event Storming to UML Generator application
-- Updated to address UUID ordering issues and 2038 timestamp problem

-- Users table
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    seq_id INT AUTO_INCREMENT UNIQUE,
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    is_active BOOLEAN DEFAULT TRUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL DEFAULT NULL,
    INDEX idx_users_email (email),
    INDEX idx_users_username (username),
    INDEX idx_users_created_at (created_at),
    INDEX idx_users_deleted_at (deleted_at)
);

-- User sessions table
CREATE TABLE IF NOT EXISTS user_sessions (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    token VARCHAR(255) UNIQUE NOT NULL,
    expires_at DATETIME NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    INDEX idx_user_sessions_user_id (user_id),
    INDEX idx_user_sessions_token (token),
    INDEX idx_user_sessions_expires_at (expires_at)
);

-- Projects table
CREATE TABLE IF NOT EXISTS projects (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    seq_id INT AUTO_INCREMENT UNIQUE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    business_process_text LONGTEXT,
    owner_user_id VARCHAR(36) NOT NULL,
    is_template BOOLEAN DEFAULT FALSE,
    miro_board_id VARCHAR(255) NULL,
    miro_export_status ENUM('pending', 'processing', 'completed', 'failed') DEFAULT 'pending',
    last_miro_export_at DATETIME NULL DEFAULT NULL,
    ai_processing_status ENUM('pending', 'processing', 'completed', 'failed') DEFAULT 'pending',
    ai_confidence_score DECIMAL(5,4) NULL DEFAULT NULL,
    last_processed_at DATETIME NULL DEFAULT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL DEFAULT NULL,
    FULLTEXT idx_projects_business_process (business_process_text),
    FULLTEXT idx_projects_name_description (name, description),
    INDEX idx_projects_owner_user_id (owner_user_id),
    INDEX idx_projects_created_at (created_at),
    INDEX idx_projects_deleted_at (deleted_at),
    INDEX idx_projects_is_template (is_template),
    INDEX idx_projects_miro_export_status (miro_export_status),
    INDEX idx_projects_ai_processing_status (ai_processing_status),
    INDEX idx_projects_user_created (owner_user_id, created_at),
    FOREIGN KEY (owner_user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Project users table (many-to-many relationship with roles)
CREATE TABLE IF NOT EXISTS project_users (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    seq_id INT AUTO_INCREMENT UNIQUE,
    project_id VARCHAR(36) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    role ENUM('viewer', 'editor') NOT NULL DEFAULT 'viewer',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL DEFAULT NULL,
    UNIQUE KEY uk_project_user (project_id, user_id),
    INDEX idx_project_users_project_id (project_id),
    INDEX idx_project_users_user_id (user_id),
    INDEX idx_project_users_role (role),
    INDEX idx_project_users_deleted_at (deleted_at),
    INDEX idx_project_users_user_created (user_id, created_at),
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Sticky notes table
CREATE TABLE IF NOT EXISTS sticky_notes (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    seq_id INT AUTO_INCREMENT UNIQUE,
    project_id VARCHAR(36) NOT NULL,
    type ENUM('domain_event', 'command', 'policy', 'actor') NOT NULL,
    content TEXT NOT NULL,
    x_position INT DEFAULT 0,
    y_position INT DEFAULT 0,
    width INT DEFAULT 200,
    height INT DEFAULT 100,
    color VARCHAR(7) NULL DEFAULT NULL,
    is_ai_generated BOOLEAN DEFAULT TRUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL DEFAULT NULL,
    INDEX idx_sticky_notes_project_id (project_id),
    INDEX idx_sticky_notes_type (type),
    INDEX idx_sticky_notes_position (project_id, x_position, y_position),
    INDEX idx_sticky_notes_created_at (created_at),
    INDEX idx_sticky_notes_deleted_at (deleted_at),
    INDEX idx_sticky_notes_project_created (project_id, created_at),
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
);

-- Note relationships table
CREATE TABLE IF NOT EXISTS note_relationships (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    seq_id INT AUTO_INCREMENT UNIQUE,
    project_id VARCHAR(36) NOT NULL,
    source_note_id VARCHAR(36) NOT NULL,
    target_note_id VARCHAR(36) NOT NULL,
    relationship_type ENUM('triggers', 'implements', 'associated_with') NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL DEFAULT NULL,
    INDEX idx_note_relationships_project_id (project_id),
    INDEX idx_note_relationships_source_note (source_note_id),
    INDEX idx_note_relationships_target_note (target_note_id),
    INDEX idx_note_relationships_type (relationship_type),
    INDEX idx_note_relationships_deleted_at (deleted_at),
    INDEX idx_note_relationships_project_created (project_id, created_at),
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE,
    FOREIGN KEY (source_note_id) REFERENCES sticky_notes(id) ON DELETE CASCADE,
    FOREIGN KEY (target_note_id) REFERENCES sticky_notes(id) ON DELETE CASCADE
);

-- UML diagrams table
CREATE TABLE IF NOT EXISTS uml_diagrams (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    seq_id INT AUTO_INCREMENT UNIQUE,
    project_id VARCHAR(36) NOT NULL,
    diagram_type ENUM('use_case', 'sequence', 'class') NOT NULL,
    svg_content LONGTEXT,
    plantuml_text LONGTEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL DEFAULT NULL,
    INDEX idx_uml_diagrams_project_id (project_id),
    INDEX idx_uml_diagrams_type (diagram_type),
    INDEX idx_uml_diagrams_created_at (created_at),
    INDEX idx_uml_diagrams_deleted_at (deleted_at),
    INDEX idx_uml_diagrams_project_created (project_id, created_at),
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
);

-- Notifications table
CREATE TABLE IF NOT EXISTS notifications (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    seq_id INT AUTO_INCREMENT UNIQUE,
    user_id VARCHAR(36) NOT NULL,
    project_id VARCHAR(36) NULL,
    message TEXT NOT NULL,
    is_read BOOLEAN DEFAULT FALSE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    read_at DATETIME NULL DEFAULT NULL,
    deleted_at DATETIME NULL DEFAULT NULL,
    INDEX idx_notifications_user_id (user_id),
    INDEX idx_notifications_project_id (project_id),
    INDEX idx_notifications_is_read (is_read),
    INDEX idx_notifications_created_at (created_at),
    INDEX idx_notifications_deleted_at (deleted_at),
    INDEX idx_notifications_user_created (user_id, created_at),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
);

-- AI processing logs table
CREATE TABLE IF NOT EXISTS ai_processing_logs (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    seq_id INT AUTO_INCREMENT UNIQUE,
    project_id VARCHAR(36) NOT NULL,
    input_text LONGTEXT,
    processing_time_ms INT,
    result_summary TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_ai_processing_logs_project_id (project_id),
    INDEX idx_ai_processing_logs_created_at (created_at),
    INDEX idx_ai_processing_logs_project_created (project_id, created_at),
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
);

-- Project invitations table
CREATE TABLE IF NOT EXISTS project_invitations (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    seq_id INT AUTO_INCREMENT UNIQUE,
    inviter_user_id VARCHAR(36) NOT NULL,
    invitee_email VARCHAR(255) NOT NULL,
    project_id VARCHAR(36) NOT NULL,
    invitation_status ENUM('pending', 'accepted', 'declined', 'expired') DEFAULT 'pending',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    expires_at DATETIME NOT NULL,
    INDEX idx_project_invitations_inviter (inviter_user_id),
    INDEX idx_project_invitations_invitee (invitee_email),
    INDEX idx_project_invitations_project (project_id),
    INDEX idx_project_invitations_status (invitation_status),
    INDEX idx_project_invitations_created (inviter_user_id, created_at),
    FOREIGN KEY (inviter_user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
);

-- User favorites table
CREATE TABLE IF NOT EXISTS user_favorites (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    seq_id INT AUTO_INCREMENT UNIQUE,
    user_id VARCHAR(36) NOT NULL,
    project_id VARCHAR(36) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uk_user_project_favorite (user_id, project_id),
    INDEX idx_user_favorites_user_id (user_id),
    INDEX idx_user_favorites_project_id (project_id),
    INDEX idx_user_favorites_created (user_id, created_at),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
);
